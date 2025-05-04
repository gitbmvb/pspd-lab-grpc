import grpc
from concurrent import futures
import psycopg2
from google.protobuf import empty_pb2
from google.protobuf.timestamp_pb2 import Timestamp
import grpc_services.service_pb2 as service_pb2
import grpc_services.service_pb2_grpc as service_pb2_grpc
from datetime import datetime

class DataService(service_pb2_grpc.DataServiceServicer):
    def __init__(self):
        self.conn = psycopg2.connect(
            dbname="pspdlabs",
            user="admin",
            password="admin",
            host="localhost",
            port="5432"
        )

    # ========== User CRUD ==========

    def LoginUser(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received LoginUser request - email: {request.email}")

                cursor.execute(
                    "SELECT 1 FROM \"USER\" WHERE email = %s AND \"password\" = %s LIMIT 1",
                    (request.email, request.password)
                )
                result = cursor.fetchone()

                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "User not found or invalid credentials")
                return empty_pb2.Empty()
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def CreateUser(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateUser request - email: {request.email}, name: {request.name}")

                cursor.execute(
                    """INSERT INTO "USER" (email, "name", "password")
                    VALUES (%s, %s, %s)
                    RETURNING email, "name", "password\"""",
                    (request.email, request.name, request.password)
                )
                result = cursor.fetchone()
                self.conn.commit()
                return service_pb2.User(
                    email=result[0],
                    name=result[1],
                    password=result[2]
                )
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.ALREADY_EXISTS if "unique" in str(e).lower() 
                              else grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateUser(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateUser request - email: {request.email}, name: {request.name}")

                cursor.execute(
                    """UPDATE "USER" 
                    SET "name" = COALESCE(%s, name), 
                        "password" = COALESCE(%s, password)
                    WHERE email = %s
                    RETURNING email, "name", \"password\"""",
                    (request.name, request.password, request.email)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
                return service_pb2.User(email=result[0], name=result[1], password=result[2])
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteUser(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteUser request - email: {request.email}")

                cursor.execute(
                    "DELETE FROM \"USER\" WHERE email = %s RETURNING email",
                    (request.email,)
                )
                if not cursor.fetchone():
                    context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
                self.conn.commit()
                return empty_pb2.Empty()
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetUser(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetUser request - email: {request.email}")

                cursor.execute(
                    "SELECT \"name\", email, \"password\" FROM \"USER\" WHERE email = %s",
                    (request.email,)
                )
                result = cursor.fetchone()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
                return service_pb2.User(name=result[0], email=result[1], password=result[2])
            except grpc.RpcError as e:
                if e.code() == grpc.StatusCode.NOT_FOUND:
                    print("User not found")
                else:
                    print(f"Other error: {e.details()}")

    def ListUsers(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print("gRPC|db - Received ListUsers request")

                cursor.execute("SELECT \"name\", email, \"password\" FROM \"USER\"")
                response = service_pb2.Users()
                for row in cursor:
                    response.users.add(name=row[0], email=row[1], password=row[2])
                return response
            except Exception as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Error listing users: {str(e)}")

    # ========== Chat CRUD ==========

    def CreateChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateChat request - subject: {request.subject}, email: {request.email}")

                cursor.execute(
                    """INSERT INTO CHAT ("subject", email)
                    VALUES (%s, %s)
                    RETURNING idChat, "subject", startDate, email""",
                    (request.subject, request.email)
                )
                result = cursor.fetchone()
                self.conn.commit()
                chat = service_pb2.Chat(idChat=result[0], subject=result[1], email=result[3])
                chat.startDate.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateChat request - id: {request.idChat}, subject: {request.subject}, email: {request.email}")

                cursor.execute(
                    """UPDATE CHAT 
                    SET "subject" = COALESCE(%s, "subject")
                    WHERE idChat = %s
                    RETURNING idChat, "subject", startDate, email""",
                    (request.subject, request.idChat)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Chat not found or access denied")
                chat = service_pb2.Chat(idChat=result[0], subject=result[1], email=result[3])
                chat.startDate.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteChat request - id: {request.id}")

                cursor.execute(
                    "DELETE FROM CHAT WHERE idChat = %s RETURNING idChat",
                    (request.idChat,)
                )
                if not cursor.fetchone():
                    context.abort(grpc.StatusCode.NOT_FOUND, "Chat not found")
                self.conn.commit()
                return empty_pb2.Empty()
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetChat request - idChat: {request.idChat}, email: {request.email}")

                if not request.idChat or not request.email:
                    context.abort(grpc.StatusCode.INVALID_ARGUMENT, "Both id and email must be provided")

                cursor.execute(
                    "SELECT idChat, \"subject\", startDate, email FROM CHAT WHERE idChat = %s AND email = %s",
                    (request.idChat, request.email)
                )
                result = cursor.fetchone()

                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Chat not found or access denied")

                chat = service_pb2.Chat(idChat=result[0], subject=result[1], email=result[3])
                chat.startDate.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")


    def ListChats(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetChats request - email: {request.email}")

                cursor.execute(
                    "SELECT idChat, \"subject\", startDate, email FROM CHAT WHERE email = %s ORDER BY startDate ASC",
                    (request.email,)
                )
                
                response = service_pb2.Chats()
                for row in cursor:
                    chat = response.chats.add(
                        idChat=row[0],
                        subject=row[1],
                        email=row[3]
                    )
                    if row[2]:  # Handle timestamp conversion if not None
                        chat.startDate.FromDatetime(row[2])
                return response
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    # ========== Message CRUD ==========

    def CreateMessage(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateMessage request - question: {request.question}, answer: {request.answer}, idChat: {request.idChat}")

                cursor.execute(
                    """INSERT INTO MESSAGE (content, idChat)
                    VALUES (%s, %s, %s)
                    RETURNING idMessage, content, "dateTime", idChat""",
                    (request.content, request.idChat)
                )
                result = cursor.fetchone()
                self.conn.commit()
                message = service_pb2.Message(
                    idMessage=result[0],
                    content=result[1],
                    idChat=result[3]
                )
                message.dateTime.FromDatetime(result[2])
                return message
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateMessage(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateMessage request - idMessage: {request.idMessage}, question: {request.question}, answer: {request.answer}")

                cursor.execute(
                    """UPDATE MESSAGE 
                    SET content = COALESCE(%s, content)
                    WHERE idMessage = %s
                    RETURNING idMessage, content, "dateTime", idChat""",
                    (request.content, request.idMessage)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Message not found")
                message = service_pb2.Message(
                    idMessage=result[0],
                    content=result[1],
                    idChat=result[3]
                )
                message.dateTime.FromDatetime(result[2])
                return message
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteMessage(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteMessage request - idMessage: {request.idMessage}")

                cursor.execute(
                    "DELETE FROM MESSAGE WHERE idMessage = %s RETURNING idMessage",
                    (request.idMessage,)
                )
                if not cursor.fetchone():
                    context.abort(grpc.StatusCode.NOT_FOUND, "Message not found")
                self.conn.commit()
                return empty_pb2.Empty()
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetMessage(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetMessage request - idMessage: {request.idMessage}, idChat: {request.idChat}")
                
                cursor.execute(
                    """SELECT m.idMessage, m.content, m."dateTime", m.idChat 
                    FROM MESSAGE m 
                    WHERE m.idMessage = %s AND m.idChat = %s""",
                    (request.idMessage, request.idChat)
                )
                result = cursor.fetchone()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, 
                                f"Message not found with idMessage={request.idMessage} in chat {request.idChat}")
                
                message = service_pb2.Message(
                    idMessage=result[0],
                    content=result[1],
                    idChat=result[3]
                )
                if result[2]:  # Handle dateTime conversion
                    message.dateTime.FromDatetime(result[2])
                return message
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def ListMessages(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetMessages request - idChat: {request.idChat}")
                
                cursor.execute(
                    """SELECT m.idMessage, m.content, m."dateTime", m.idChat 
                    FROM MESSAGE m 
                    WHERE m.idChat = %s 
                    ORDER BY m."dateTime" DESC""",
                    (request.idChat,)
                )

                response = service_pb2.Messages()
                
                for row in cursor:
                    message = response.messages.add(
                        idMessage=row[0],
                        content=row[1],
                        idChat=row[3]
                    )
                    if row[2]:
                        message.dateTime.FromDatetime(row[2])
                
                return response
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_DataServiceServicer_to_server(DataService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started on port 50051")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()