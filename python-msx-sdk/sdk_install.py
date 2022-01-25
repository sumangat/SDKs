import python_msx_sdk
from python_msx_sdk.api_client import ApiClient
from python_msx_sdk.apis import SecurityApi, TenantsApi
from python_msx_sdk.models import TenantCreate
import uuid
import base64
import urllib3
urllib3.disable_warnings()

MY_SERVER_URL = "https://dev-plt-aio4.lab.ciscomsx.com"
VERIFY_SSL = False  # TODO - Do not defeat the SSL certificate in production code!
MY_CLIENT_ID = "hello-world-service-private-client"
MY_CLIENT_SECRET = "make-up-a-private-client-secret-and-keep-it-safe"
MY_USERNAME = "superuser"
MY_PASSWORD = "GLCpiJ9GSLp_?Lq#ATxqlf~9Up)%Wl7@"


def get_api_client(access_token):
    configuration = python_msx_sdk.Configuration(MY_SERVER_URL)
    configuration.verify_ssl = VERIFY_SSL
    api_client = ApiClient(configuration=configuration)
    if access_token:
        api_client.set_default_header("Authorization", "Bearer " + access_token)
    return api_client


def get_access_token(authorization, username, password):
    api_client = get_api_client(access_token=None)
    security_api = SecurityApi(api_client)
    response = security_api.get_access_token(
        authorization=authorization,
        grant_type="password",
        username=username,
        password=password)

    return response.access_token


def format_tenant(x):
    return x.id + ", " + x.name + "\n" + x.description + "\n"


def main():
    basic_token = base64.b64encode(str.encode(MY_CLIENT_ID + ":" + MY_CLIENT_SECRET))
    basic_authorization = "Basic " + basic_token.decode()

    access_token = get_access_token(authorization=basic_authorization, username=MY_USERNAME, password=MY_PASSWORD)
    api_client = get_api_client(access_token=access_token)
    tenants_api = TenantsApi(api_client)

    body = TenantCreate(
        name="my_unique_tenant_name_" + str(uuid.uuid4()),
        description="Monsters exist, but they are too few in number to be truly dangerous. More dangerous are the common men, the functionaries ready to believe and to act without asking questions. --Primo Levi",
        email= "email@cisco.com",
        url="https://cisco.com")

    tenant1 = tenants_api.create_tenant(body)
    print("Tenant Identifier: " + tenant1.id)

    page0 = tenants_api.get_tenants_page(0, 100)
    tenant_id_list = [d['id'] for d in page0['contents']]

    # Keep it for debugging
    # print("\n".join(map(lambda x: format_tenant(x), page0.contents)))

    if tenant1.id not in tenant_id_list:
        raise Exception(f'Failed to find tenant with id {tenant1.id}')

    res = tenants_api.delete_tenant(tenant1.id)
    print('Tenant got deleted:'  + tenant1.id)

   

if __name__ == "__main__":
    main()