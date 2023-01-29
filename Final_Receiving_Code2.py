import socket                   # Import socket module
from cryptography.fernet import Fernet
#port = 49534                 # Reserve a port for your service every new transfer wants a new port or you must wait.
#port = 49534 
port = 53734  #:53734        LAPTOP-S8QBTLN6:53733
s = socket.socket()             # Create a socket object
host = socket.gethostname()#"127.0.0.1"   # Get local machine name
s.bind((host, port))            # Bind to the port
print('Host: ' ,host)
print('Port: ' ,port)
s.listen(5)                     # Now wait for client connection.

#key = Fernet.generate_key()
key = b'f-Fd2vxmXx_3xO6Y1vdZ0eCnRN0ocBJ_cfpxC8W2JII='
fernet = Fernet(key)
print('Server listening....')

while True:
    conn, addr = s.accept()     # Establish connection with client.
    print('Got connection from ', addr)
    if (addr[0] == '192.168.43.72'):
        with open('Outfile_Instance1.json', 'a') as f:
            print('file opened')
            while True:
                print('receiving data...')
                data = conn.recv(1024)
                print('Data Data',data)         
                print('Print Data...',data)
                if not data:
                    break
                # write data to a file
                f.write(data.decode())
        f.close()
    if (addr[0] == '192.168.43.46'):
        with open('Outfile_Instance2.json', 'a') as f:
            print('file opened')
            while True:
                print('receiving data...')
                data = conn.recv(1024)
                print('Print Data...',data)
                print('data=%s', (data))
                if not data:
                    break
                # write data to a file
                f.write(data.decode())
        f.close()
    
    if (addr[0] == '192.168.43.167'):
        with open('Outfile_Instance3.json', 'a') as f:
            print('file opened')
            while True:
                print('receiving data...')
                data = conn.recv(1024)
                print('Print Data...',data)
                print('data=%s', (data))
                if not data:
                    break
                # write data to a file
                f.write(data.decode())
        f.close()      
        


print('Successfully get the file')
print('connection closed') 