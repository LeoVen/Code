import * as jose from 'jose';
import * as crypto from 'crypto';
import fs from 'node:fs';
import util from 'util';
const exec = util.promisify(require('child_process').exec);

async function main(): Promise<void> {
    if (process.env.NO_SETUP !== 'true') {
        await exec('chmod 744 ./gen_keys.sh');
        await exec('./gen_keys.sh');
    } else {
        console.log('Skipping setup\n');
    }

    const private_key_bytes = fs.readFileSync('jwtRS256.key');
    const public_key_bytes = fs.readFileSync('jwtRS256.key.pub');

    const private_key = crypto.createPrivateKey(private_key_bytes);
    const public_key = crypto.createPublicKey(public_key_bytes);

    const signed = await new jose.SignJWT()
        .setExpirationTime('2 hours')
        .setIssuedAt(new Date())
        .setIssuer('test app')
        .setProtectedHeader({
            alg: 'RS256',
            user: 'my user',
        })
        .sign(private_key);

    const jwk = public_key.export({ format: 'jwk' });

    console.log(JSON.stringify(jwk) + '\n');
    console.log(signed + '\n');

    try {
        const result = await jose.jwtVerify(signed, await jose.importJWK(jwk));
        console.log(result);
    } catch (e) {
        console.log(e);
    }
}

main();
