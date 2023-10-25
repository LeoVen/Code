# run example with postgresql

npx prisma format --schema ./prisma/postgre.prisma
npx prisma generate --schema ./prisma/postgre.prisma
npm run postgre

# generate the initdb sql
# needs to have postgresql running
npx prisma migrate dev --create-only --schema ./prisma/postgre.prisma
