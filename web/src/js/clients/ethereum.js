import ed25519 from 'supercop.js';
import generatePassword from 'password-generator';
import Web3 from 'web3';

export default class {
    constructor() {
        const provider = Web3.currentProvider || 'http://0.0.0.0:8545';
        this.web3 = new Web3(provider);
    }

    createAccount() {
        return new Promise((resolve, reject) => {
            const pair = ed25519.createKeyPair(ed25519.createSeed());
            const password = generatePassword(32);

            this.web3.eth.personal.newAccount(password, (err, address) => {
                if (err) {
                    reject(err);
                }

                resolve({
                    address: address,
                    pubkey: pair.publicKey,
                    privkey: pair.secretKey
                });
            });
        });
    }
}
