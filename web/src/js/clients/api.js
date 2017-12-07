export default class {
    constructor() {
        this.opts = {
            mode: 'cors',
            method: 'POST'
        };
    }

    upload(file) {
        return fetch(
            'http://localhost:8080/v1/files',
            Object.assign(this.opts, { body: file })
        ).then(res => {
            console.log('done: ', res);
        });
    }

    createAccount() {
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
    }
}
