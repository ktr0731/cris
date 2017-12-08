import Web3 from 'web3';

import Cris from '../../../contracts/Cris.sol';
import ed25519 from 'supercop.js';
import generatePassword from 'password-generator';

export default class {
    constructor() {
        console.log('init eth client');
        const provider = new Web3.providers.HttpProvider(
            'http://localhost:8545'
        );

        Cris.setProvider(provider);
        this.cris = Cris.deployed();
        this.web3 = new Web3(provider);
    }

    createAccount() {
        return new Promise((resolve, reject) => {
            console.log('create new account');
            const pair = ed25519.createKeyPair(ed25519.createSeed());

            const password = generatePassword(32);
            localStorage['password'] = password;

            this.web3.personal.newAccount(password, (err, address) => {
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

    unlock() {
        console.log(this.web3);
        this.web3.personal.unlockAccount(
            localStorage['address'],
            localStorage['password'],
            24 * 60 * 60,
            err => {
                if (!err) {
                    console.log(err);
                }
                console.log('unlocked');
            }
        );
    }

    store(hash) {
        this.unlock();
        return this.cris.store(hash, { from: localStorage['address'] });
    }
}
