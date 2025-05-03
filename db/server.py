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

    def CreateUser(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateUser request - email: {request.email}, name: {request.name}")

                cursor.execute(
                    """INSERT INTO users (email, name, password)
                    VALUES (%s, %s, %s)
                    RETURNING email, name, password""",
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
                    """UPDATE users 
                    SET name = COALESCE(%s, name), 
                        password = COALESCE(%s, password)
                    WHERE email = %s
                    RETURNING email, name, password""",
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
                    "DELETE FROM users WHERE email = %s RETURNING email",
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
                    "SELECT name, email, password FROM users WHERE email = %s",
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

                cursor.execute("SELECT name, email, password FROM users")
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
                print(f"gRPC|db - Received CreateChat request - topic: {request.topic}, email: {request.email}")

                cursor.execute(
                    """INSERT INTO chats (topic, email)
                    VALUES (%s, %s)
                    RETURNING id, topic, created_at, email""",
                    (request.topic, request.email)
                )
                result = cursor.fetchone()
                self.conn.commit()
                chat = service_pb2.Chat(id=result[0], topic=result[1], email=result[3])
                chat.created_at.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateChat request - id: {request.id}, topic: {request.topic}, email: {request.email}")

                cursor.execute(
                    """UPDATE chats 
                    SET topic = COALESCE(%s, topic)
                    WHERE id = %s AND email = %s
                    RETURNING id, topic, created_at, email""",
                    (request.topic, request.id, request.email)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Chat not found or access denied")
                chat = service_pb2.Chat(id=result[0], topic=result[1], email=result[3])
                chat.created_at.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteChat request - id: {request.id}")

                cursor.execute(
                    "DELETE FROM chats WHERE id = %s RETURNING id",
                    (request.id,)
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
                print(f"gRPC|db - Received GetChat request - id: {request.id}, email: {request.email}")

                if not request.id or not request.email:
                    context.abort(grpc.StatusCode.INVALID_ARGUMENT, "Both id and email must be provided")

                cursor.execute(
                    "SELECT id, topic, created_at, email FROM chats WHERE id = %s AND email = %s",
                    (request.id, request.email)
                )
                result = cursor.fetchone()

                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Chat not found or access denied")

                chat = service_pb2.Chat(id=result[0], topic=result[1], email=result[3])
                chat.created_at.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")


    def GetChats(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetChats request - email: {request.email}")

                cursor.execute(
                    "SELECT id, topic, created_at, email FROM chats WHERE email = %s ORDER BY created_at ASC",
                    (request.email,)
                )
                
                response = service_pb2.Chats()
                for row in cursor:
                    chat = response.chats.add(
                        id=row[0],
                        topic=row[1],
                        email=row[3]
                    )
                    if row[2]:  # Handle timestamp conversion if not None
                        chat.created_at.FromDatetime(row[2])
                return response
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    # ========== Message CRUD ==========

    def CreateMessage(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateMessage request - question: {request.question}, answer: {request.answer}, chat_id: {request.chat_id}")

                cursor.execute(
                    """INSERT INTO messages (question, answer, chat_id)
                    VALUES (%s, %s, %s)
                    RETURNING message_id, question, answer, timestamp, chat_id""",
                    (request.question, request.answer, request.chat_id)
                )
                result = cursor.fetchone()
                self.conn.commit()
                message = service_pb2.Message(
                    message_id=result[0],
                    question=result[1],
                    answer=result[2],
                    chat_id=result[4]
                )
                message.timestamp.FromDatetime(result[3])
                return message
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateMessage(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateMessage request - message_id: {request.message_id}, question: {request.question}, answer: {request.answer}")

                cursor.execute(
                    """UPDATE messages 
                    SET question = COALESCE(%s, question),
                        answer = COALESCE(%s, answer)
                    WHERE message_id = %s
                    RETURNING message_id, question, answer, timestamp, chat_id""",
                    (request.question, request.answer, request.message_id)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Message not found")
                message = service_pb2.Message(
                    message_id=result[0],
                    question=result[1],
                    answer=result[2],
                    chat_id=result[4]
                )
                message.timestamp.FromDatetime(result[3])
                return message
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteMessage(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteMessage request - message_id: {request.message_id}")

                cursor.execute(
                    "DELETE FROM messages WHERE message_id = %s RETURNING message_id",
                    (request.message_id,)
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
                print(f"gRPC|db - Received GetMessage request - message_id: {request.message_id}, chat_id: {request.chat_id}")
                
                cursor.execute(
                    """SELECT m.message_id, m.question, m.answer, m.timestamp, m.chat_id 
                    FROM messages m 
                    WHERE m.message_id = %s AND m.chat_id = %s""",
                    (request.message_id, request.chat_id)
                )
                result = cursor.fetchone()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, 
                                f"Message not found with message_id={request.message_id} in chat {request.chat_id}")
                
                message = service_pb2.Message(
                    message_id=result[0],
                    question=result[1],
                    answer=result[2],
                    chat_id=result[4]
                )
                if result[3]:  # Handle timestamp conversion
                    message.timestamp.FromDatetime(result[3])
                return message
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetMessages(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetMessages request - chat_id: {request.chat_id}")
                
                cursor.execute(
                    """SELECT m.message_id, m.question, m.answer, m.timestamp, m.chat_id 
                    FROM messages m 
                    WHERE m.chat_id = %s 
                    ORDER BY m.timestamp DESC""",
                    (request.chat_id,)
                )

                response = service_pb2.Messages()
                
                for row in cursor:
                    message = response.messages.add(
                        message_id=row[0],
                        question=row[1],
                        answer=row[2],
                        chat_id=row[4]
                    )
                    if row[3]:
                        message.timestamp.FromDatetime(row[3])
                
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