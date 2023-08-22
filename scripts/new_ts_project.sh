# change projectName
mkdir projectName && cd $_

npm init --yes
npm i -D typescript ts-node

./node_modules/.bin/tsc --init
echo 'console.log("Hello, World!")' > index.ts

# put this in scripts
node --loader ts-node/esm index.ts

npx tsc --init
