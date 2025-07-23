import socket

HOST = 'localhost'
PORT = 8080

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
	s.bind((HOST, PORT))
	s.listen()
	conn, addr = s.accept()
	while True:
		conn, addr = s.accept()
		with conn:
			print(f"Брух подключен: {addr}")
			conn.sendall(b'OK\n')
