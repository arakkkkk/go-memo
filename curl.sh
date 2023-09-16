# curl -b cookie.txt --request POST localhost:8080/api/v1/memo/create \
# 	--header 'content-type: application/json' \
# 	--data '{"description": "test todo"}'

# curl -b cookie.txt --request GET localhost:8080/api/v1/memo


# curl --request GET localhost:8080/api/v1/user/list
# curl -c cookie.txt -b cookie.txt --request GET localhost:8080/api/v1/user/logout

# curl --request POST localhost:8080/api/v1/user/register \
# 	--header 'content-type: application/json' \
# 	--data '{"name": "arakkk", "password": "pass", "password_confirm": "pass"}'

# curl -c cookie.txt --request POST localhost:8080/api/v1/user/login \
# 	--header 'content-type: application/json' \
# 	--data '{"name": "arakkk", "password": "pass"}'
