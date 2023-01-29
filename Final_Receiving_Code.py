import socket                   # Import socket module
from cryptography.fernet import Fernet
port = 49534                 # Reserve a port for your service every new transfer wants a new port or you must wait.
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
    if (addr[0] == '192.168.0.206'):
        data=conn.recv(10000)
        with open('Outfile_nameOther.json', 'a') as f:
            print('file opened')
            while True:
                print('receiving data...')
                data = conn.recv(1024)
                if (data == b''):
                    break
                print('Print Data...',data)
                decrypted_data = fernet.decrypt(data)
                decrypted_data=decrypted_data.decode()
                print('data=%s', (decrypted_data))
                if not decrypted_data:
                    break
                # write data to a file
                f.write(decrypted_data)
        f.close()
    if (addr[0] == '192.168.0.59'):
        data=conn.recv(10000)
        with open('Outfile_name2MyName.json', 'a') as f:
            print('file opened')
            while True:
                print('receiving data...')
                data = conn.recv(1024)
                if (data == b''):
                    break
                print('Print Data...',data)
                decrypted_data = fernet.decrypt(data)
                decrypted_data=decrypted_data.decode()
                print('data=%s', (decrypted_data))
                if not decrypted_data:
                    break
                # write data to a file
                f.write(decrypted_data)
        f.close()
        


print('Successfully get the file')
#s.close()
#conn.close()
print('connection closed') 