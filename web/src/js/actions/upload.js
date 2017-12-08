import APIClient from '../clients/api';
import EthClient from '../clients/ethereum';
import ed25519 from 'supercop.js';

export default (file, store) => {
    // TODO: send signature also
    let hash = null;
    let token = null;
    return new Promise((resolve, _) => {
        const reader = new FileReader();
        reader.onload = e => {
            const contentBuf = new Buffer(reader.result.length);
            const pubkeyBuf = new Buffer(localStorage['pubkey'].length);
            const privkeyBuf = new Buffer(localStorage['privkey'].length);
            contentBuf.fill(reader.result);
            pubkeyBuf.fill(localStorage['pubkey']);
            privkeyBuf.fill(localStorage['privkey']);
            const encryptedContent = ed25519.sign(
                contentBuf,
                pubkeyBuf,
                privkeyBuf
            );
            hash = Buffer.from(encryptedContent).toString('hex');

            return resolve(encryptedContent);
        };
        reader.readAsText(file);
    })
        .then(new APIClient().upload)
        .then(res => res.json())
        .then(res => {
            token = res.token;
            return new EthClient().store(hash);
        })
        .then(hash => {
            console.log(hash);
            store.addUploadedFile({
                name: file.name,
                hash: hash,
                token: token,
                txHash: hash,
                date: file.lastModifiedDate
            });
        });
};
