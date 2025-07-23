import socket

HOST = 'localhost'
PORT = 8080

try:
	with socket.create_connection((HOST, PORT)) as sock:
		response = sock.recv(1024).decode('utf-8')
		if response == "OK\n":
			print("брух, все гуд")
		else:
			print("брух, чет не то, это ваще не окей")
			print(f"ответ: {response}")
except ConnectionRefusedError:
	print("Это что, брух, что с портом")
except Exception as e:
	print("Это что, брух. принять не смогли")
	print(e)
