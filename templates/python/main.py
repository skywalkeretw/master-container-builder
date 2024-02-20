import threading
import http_server as httpServer
import messaging as messaging
import os

def getenv_bool(key, default_value=False):
    env_value = os.getenv(key)
    
    if env_value is None:
        return default_value
    
    try:
        # Attempt to parse the environment variable as a boolean
        return bool(int(env_value))
    except ValueError:
        # If parsing fails, return the default value
        return default_value

if __name__ == '__main__':
    # Run handler1 on port 8080
    if getenv_bool("HTTP", True):
        thread1 = threading.Thread(target=httpServer.run_server)
        thread1.start()

    # Run handler2 on port 8081
    if getenv_bool("MESSAGING", False):
        thread2 = threading.Thread(target=messaging.listen_to_rabbitmq)
        thread2.start()
