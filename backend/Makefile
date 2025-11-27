prisma-gen:
	go run github.com/steebchen/prisma-client-go generate --schema=./internal/db/schema.prisma

prisma-push:
	go run github.com/steebchen/prisma-client-go db push --schema=./internal/db/schema.prisma
 
open-db:
	docker exec -it nexus-db psql -U nexus -d nexus
#go run github.com/steebchen/prisma-client-go db pull --schema=./internal/db/schema.prisma
 
#go run github.com/steebchen/prisma-client-go migrate dev --schema=./internal/db/schema.prisma
 
#go run github.com/steebchen/prisma-client-go migrate deploy --schema=./internal/db/schema.prisma