# run example with mongodb

npx prisma format --schema ./prisma/mongo.prisma
npx prisma generate --schema ./prisma/mongo.prisma
npm run mongo
