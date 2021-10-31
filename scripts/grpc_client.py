import asyncio

from grpclib.client import Channel

from ozonmp.srv_verification_api.v1.srv_verification_api_grpc import SrvVerificationApiServiceStub
from ozonmp.srv_verification_api.v1.srv_verification_api_pb2 import DescribeVerificationV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = SrvVerificationApiServiceStub(channel)

        req = DescribeVerificationV1Request(verification_id=1)
        reply = await client.DescribeVerificationV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
