import 'riot';

import '../tags/app.tag';
import '../tags/upload.tag';

import Web3 from 'web3';
import ed25519 from 'ed25519-supercop';

const provider = Web3.currentProvider || 'http://localhost:8545';
const web3 = new Web3(provider);

const init = () => {
    if (
        !localStorage['address'] ||
        !localStorage['pubkey'] ||
        !localStorage['privkey']
    ) {
        createAccount().then;
    }

    riot.mount('*');
};

const createAccount = () => {
    return new Promise((resolve, reject) => {
        const pair = ed25519.createKeyPair(ed25519.createSeed());
        web3.eth.personal.newAccount(pair.secretKey, (err, address) => {
            if (err) {
                reject(err);
            }

            localStorage['address'] = address;
            localStorage['pubkey'] = pair.pubkey;
            localStorage['privkey'] = pair.secretKey;
            resolve();
        });
    });
};

init();
