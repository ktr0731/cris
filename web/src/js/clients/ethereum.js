import ed25519 from 'supercop.js';
import generatePassword from 'password-generator';
import Web3 from 'web3';

import Cris from '../../../contracts/Cris.sol';

export default class {
    constructor() {
        const provider =
            Web3.currentProvider ||
            new Web3.providers.HttpProvider('http://localhost:8545');
            // new Web3.providers.HttpProvider('http://35.196.45.234:8545');

        Cris.setProvider(provider);
        this.cris = Cris.deployed();
    }

    createAccount() {
        return new Promise((resolve, reject) => {
            const pair = ed25519.createKeyPair(ed25519.createSeed());

            const password = generatePassword(32);

            this.cris.eth.personal.newAccount(password, (err, address) => {
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
