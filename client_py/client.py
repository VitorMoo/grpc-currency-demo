
import grpc

import currency_pb2
import currency_pb2_grpc

def run():
	print("Attempting to connect to gRPC server at port 50051")

	channel = grpc.insecure_channel('localhost:50051')

	stub = currency_pb2_grpc.CurrencyServiceStub(channel)

	print("Connected")

	request = currency_pb2.RateRequest(
		from_currency="USD",
		to_currency="BRL"
	)

	response = stub.GetRate(request)

	print("-" * 20)
	print(f"Rate received for USD to BRL: {response.price:.2f}")
	print("-" * 20)

if __name__ == '__main__':
	run()
