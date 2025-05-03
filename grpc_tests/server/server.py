import grpc
from concurrent import futures
import psycopg2
from google.protobuf import empty_pb2
from google.protobuf.timestamp_pb2 import Timestamp
import service_pb2
import service_pb2_grpc
from datetime import datetime

class DataService(service_pb2_grpc.DataServiceServicer):
    def __init__(self):
        # Initialize database connection
        self.conn = psycopg2.connect(
            dbname="pspd",
            user="admin",
            password="admin",
            host="localhost",
            port="5432"
        )

    # ========== Usuario CRUD ==========
    
    def CreateUsuario(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateUsuario request - email: {request.email}, nome: {request.nome}")

                cursor.execute(
                    """INSERT INTO usuario (email, nome, senha)
                    VALUES (%s, %s, %s)
                    RETURNING email, nome, senha""",
                    (request.email, request.nome, request.senha)
                )
                result = cursor.fetchone()
                self.conn.commit()
                return service_pb2.Usuario(
                    email=result[0],
                    nome=result[1],
                    senha=result[2]
                )
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.ALREADY_EXISTS if "unique" in str(e).lower() 
                            else grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateUsuario(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateUsuario request - email: {request.email}, nome: {request.nome}")

                cursor.execute(
                    """UPDATE usuario 
                    SET nome = COALESCE(%s, nome), 
                        senha = COALESCE(%s, senha)
                    WHERE email = %s
                    RETURNING email, nome, senha""",
                    (request.nome, request.senha, request.email)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
                return service_pb2.Usuario(**dict(zip(('email', 'nome', 'senha'), result)))
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteUsuario(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteUsuario request - email: {request.email}")

                cursor.execute(
                    "DELETE FROM usuario WHERE email = %s RETURNING email",
                    (request.email,)
                )
                if not cursor.fetchone():
                    context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
                self.conn.commit()
                return empty_pb2.Empty()
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetUsuario(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetUsuario request - email: {request.email}")

                cursor.execute(
                    "SELECT nome, email, senha FROM usuario WHERE email = %s",
                    (request.email,)
                )
                result = cursor.fetchone()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "User not found")
                return service_pb2.Usuario(
                    nome=result[0],
                    email=result[1],
                    senha=result[2]
                )
            except grpc.RpcError as e:
                if e.code() == grpc.StatusCode.NOT_FOUND:
                    print("User not found")
                else:
                    print(f"Other error: {e.details()}")
    
    def GetUsuarios(self, request, context):
        with self.conn.cursor() as cursor:
            try:    
                print("gRPC|db - Received GetUsuarios request")

                cursor.execute("SELECT nome, email, senha FROM usuario")
                response = service_pb2.Usuarios()
                for row in cursor:
                    response.usuarios.add(
                        nome=row[0],
                        email=row[1],
                        senha=row[2]
                    )
                return response
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")
            except Exception as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Unexpected error: {str(e)}")

    # ========== Chat CRUD ==========
    
    def CreateChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateChat request - tema: {request.tema}, email: {request.email}")

                cursor.execute(
                    """INSERT INTO chat (tema, email)
                    VALUES (%s, %s)
                    RETURNING idChat, tema, dataCriacao, email""",
                    (request.tema, request.email)
                )
                result = cursor.fetchone()
                self.conn.commit()
                chat = service_pb2.Chat(
                    id_chat=result[0],
                    tema=result[1],
                    email=result[3]
                )
                chat.data_criacao.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateChat request - idChat: {request.id_chat}, tema: {request.tema}, email: {request.email}")

                cursor.execute(
                    """UPDATE chat 
                    SET tema = COALESCE(%s, tema)
                    WHERE idChat = %s AND email = %s
                    RETURNING idChat, tema, dataCriacao, email""",
                    (request.tema, request.id_chat, request.email)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Chat not found or access denied")
                chat = service_pb2.Chat(
                    id_chat=result[0],
                    tema=result[1],
                    email=result[3]
                )
                chat.data_criacao.FromDatetime(result[2])
                return chat
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteChat(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteChat request - idChat: {request.idChat}")

                cursor.execute(
                    "DELETE FROM chat WHERE idChat = %s RETURNING idChat",
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
            
                # Validate required fields
                if not request.idChat or not request.email:
                    context.abort(grpc.StatusCode.INVALID_ARGUMENT, "Both idChat and email must be provided")
                
                cursor.execute(
                    "SELECT idChat, tema, dataCriacao, email FROM CHAT WHERE idChat = %s AND email = %s",
                    (request.idChat, request.email)
                )
                result = cursor.fetchone()

                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Chat not found or access denied")
                
                chat = service_pb2.Chat(
                    id_chat=result[0],  # proto uses id_chat
                    tema=result[1],
                    email=result[3]
                )
                if result[2]:  # Handle timestamp conversion
                    chat.data_criacao.FromDatetime(result[2])
                return chat
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetChats(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received GetChats request - email: {request.email}")

                cursor.execute(
                    "SELECT idChat, tema, dataCriacao, email FROM CHAT WHERE email = %s ORDER BY dataCriacao ASC",
                    (request.email,)
                )
                
                response = service_pb2.Chats()
                for row in cursor:
                    chat = response.chats.add(
                        id_chat=row[0],
                        tema=row[1],
                        email=row[3]
                    )
                    if row[2]:  # Handle timestamp conversion if not None
                        chat.data_criacao.FromDatetime(row[2])
                return response
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    # ========== Mensagem CRUD ==========

    def CreateMensagem(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received CreateMensagem request - pergunta: {request.pergunta}, resposta: {request.resposta}, idChat: {request.idChat}")

                cursor.execute(
                    """INSERT INTO mensagem (pergunta, resposta, idChat)
                    VALUES (%s, %s, %s)
                    RETURNING idMensagem, pergunta, resposta, dataHora, idChat""",
                    (request.pergunta, request.resposta, request.idChat)
                )
                result = cursor.fetchone()
                self.conn.commit()
                mensagem = service_pb2.Mensagem(
                    idMensagem=result[0],
                    pergunta=result[1],
                    resposta=result[2],
                    idChat=result[4]
                )
                mensagem.dataHora.FromDatetime(result[3])
                return mensagem
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def UpdateMensagem(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received UpdateMensagem request - idMensagem: {request.idMensagem}, pergunta: {request.pergunta}, resposta: {request.resposta}")

                cursor.execute(
                    """UPDATE mensagem 
                    SET pergunta = COALESCE(%s, pergunta),
                        resposta = COALESCE(%s, resposta)
                    WHERE idMensagem = %s
                    RETURNING idMensagem, pergunta, resposta, dataHora, idChat""",
                    (request.pergunta, request.resposta, request.idMensagem)
                )
                result = cursor.fetchone()
                self.conn.commit()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, "Message not found")
                mensagem = service_pb2.Mensagem(
                    idMensagem=result[0],
                    pergunta=result[1],
                    resposta=result[2],
                    idChat=result[4]
                )
                mensagem.dataHora.FromDatetime(result[3])
                return mensagem
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def DeleteMensagem(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Received DeleteMensagem request - idMensagem: {request.idMensagem}")

                cursor.execute(
                    "DELETE FROM mensagem WHERE idMensagem = %s RETURNING idMensagem",
                    (request.idMensagem,)
                )
                if not cursor.fetchone():
                    context.abort(grpc.StatusCode.NOT_FOUND, "Message not found")
                self.conn.commit()
                return empty_pb2.Empty()
            except psycopg2.Error as e:
                self.conn.rollback()
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetMensagem(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Recieved GetMensagem request - idMensagem: {request.idMensagem}, idChat: {request.idChat}")
                
                cursor.execute(
                    """SELECT m.idmensagem, m.pergunta, m.resposta, m.dataHora, m.idchat 
                    FROM mensagem m 
                    WHERE m.idmensagem = %s AND m.idchat = %s""",
                    (request.idMensagem, request.idChat)
                )
                result = cursor.fetchone()
                if not result:
                    context.abort(grpc.StatusCode.NOT_FOUND, 
                                f"Message not found with idMensagem={request.idMensagem} in chat {request.idChat}")
                
                mensagem = service_pb2.Mensagem(
                    idMensagem=result[0],
                    pergunta=result[1],
                    resposta=result[2],
                    idChat=result[4]
                )
                if result[3]:  # Handle timestamp conversion
                    mensagem.dataHora.FromDatetime(result[3])
                return mensagem
                
            except psycopg2.Error as e:
                context.abort(grpc.StatusCode.INTERNAL, f"Database error: {str(e)}")

    def GetMensagens(self, request, context):
        with self.conn.cursor() as cursor:
            try:
                print(f"gRPC|db - Recieved GetMensagens request - idChat: {request.idChat}")
                
                # Convert single parameter to a tuple explicitly
                params = (request.idChat,)
                
                cursor.execute(
                    """SELECT m.idmensagem, m.pergunta, m.resposta, m.dataHora, m.idchat 
                    FROM mensagem m 
                    WHERE m.idchat = %s 
                    ORDER BY m.dataHora DESC""",
                    params
                )

                response = service_pb2.Mensagens()
                
                for row in cursor:
                    mensagem = response.mensagens.add(
                        idMensagem=row[0],
                        pergunta=row[1],
                        resposta=row[2],
                        idChat=row[4]
                    )
                    if row[3]:
                        mensagem.dataHora.FromDatetime(row[3])
                
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