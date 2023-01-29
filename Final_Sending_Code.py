from cryptography.fernet import Fernet
import socket                   

s = socket.socket()             
host = "LAPTOP-S8QBTLN6" 
#port = 21128
port = 49534

s.connect((host, port))
s.send("Hello".encode())

#key = Fernet.generate_key()
key = b'f-Fd2vxmXx_3xO6Y1vdZ0eCnRN0ocBJ_cfpxC8W2JII='
filename='C:\\Users\\himan\\.spyder-py3\\cv assignment\\Scalable_Project4\\sensor1.json' #In the same folder or path is this file running must the file you want to tranfser to be
f = open(filename,'r')
l = f.read(1024)

while (l):
    fernet = Fernet(key)
    print(key)
    encrypted = fernet.encrypt(l.encode())
    print('encrypted data ',encrypted)
    s.send(encrypted)
    print('Sent ',repr(l))
    l = f.read(1024)
    f.close()

print('Done sending')
s.close()

#import time
#while True:
#    # Code executed here
#    print("Hello World")
#    time.sleep(60)
#http://127.0.0.1:9090/battery/charge