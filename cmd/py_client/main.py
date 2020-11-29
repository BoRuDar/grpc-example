import api_pb2
import api_pb2_grpc
import grpc


def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = api_pb2_grpc.CalcStub(channel)
        response = stub.Calculate(api_pb2.Request(a=1.2, b=3.3))
    print("client received: ", response.result)


if __name__ == '__main__':
    run()
