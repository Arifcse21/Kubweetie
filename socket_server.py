import socket
import threading
import random
import string

# Store channels and connected clients
channels = {}

# Generates a random channel ID
def generate_channel_id(length=6):
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=length))

# Client handler function for relaying messages
def handle_client(client_socket, channel_id, client_id):
    while True:
        try:
            message = client_socket.recv(1024).decode()
            if message:
                print(f"Message from {client_id} in {channel_id}: {message}")
                
                # Relay message to the other client in the same channel
                for client, other_client_id in channels[channel_id]:
                    if client != client_socket:  # Don't send message back to the sender
                        client.sendall(f"{client_id}: {message}".encode())
            else:
                break
        except Exception as e:
            print(f"Error with {client_id} in {channel_id}: {e}")
            break

    # Remove client from channel on disconnect
    client_socket.close()
    channels[channel_id] = [(client, id) for client, id in channels[channel_id] if client != client_socket]
    if not channels[channel_id]:  # Delete empty channels
        del channels[channel_id]

# Main server function
def start_server(host='0.0.0.0', port=5050):  # Bind to all interfaces
    # Create the server socket
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.bind((host, port))
    server.listen()
    print(f"Server started on {host}:{port}")

    while True:
        # Accept a new client
        client_socket, client_addr = server.accept()
        print(f"Client connected from {client_addr}")

        # Get channel ID from client
        client_socket.sendall("Enter channel ID to join or press Enter to create a new channel: ".encode())
        channel_id = client_socket.recv(1024).decode().strip()

        # Generate a new channel if ID is not provided
        if not channel_id:
            channel_id = generate_channel_id()
            client_socket.sendall(f"New channel created with ID: {channel_id}".encode())
            channels[channel_id] = []

        elif channel_id not in channels:
            client_socket.sendall("Channel not found. Creating a new one.\n".encode())
            channels[channel_id] = []
        else:
            client_socket.sendall(f"Joined channel: {channel_id}".encode())

        # Assign a unique client identifier and add client to the channel
        client_id = f"Client {len(channels[channel_id]) + 1}"
        channels[channel_id].append((client_socket, client_id))
        print(f"{client_id} joined channel {channel_id}")

        # Notify other clients in the channel of the new client
        for client, _ in channels[channel_id]:
            if client != client_socket:
                client.sendall(f"{client_id} has joined the channel.".encode())

        # Start a thread to handle messages for this client
        thread = threading.Thread(target=handle_client, args=(client_socket, channel_id, client_id))
        thread.start()

# Start the server
start_server()
