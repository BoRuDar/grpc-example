import api_pb2
import api_pb2_grpc
import grpc


def run(a: float, b: float):
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = api_pb2_grpc.CalcStub(channel)
        response = stub.Calculate(api_pb2.Request(a=a, b=b))
    print("client received: ", response.result)


if __name__ == '__main__':
    run(13.4, 34.4)
