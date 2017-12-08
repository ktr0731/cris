import APIClient from '../clients/api';
import EthClient from '../clients/ethereum';
import ed25519 from 'supercop.js';
import { Base64 } from 'js-base64';

export default (file, store) => {
    // TODO: send signature also
    let hash = null;
    let token = null;
    return new Promise((resolve, _) => {
        const reader = new FileReader();
        reader.onload = e => {
            const pubkey = Base64.decode(localStorage['pubkey']);
            const privkey = Base64.decode(localStorage['privkey']);
            const contentBuf = new Buffer(reader.result.length);
            const pubkeyBuf = new Buffer(pubkey.length);
            const privkeyBuf = new Buffer(privkey.length);
            contentBuf.fill(reader.result);
            pubkeyBuf.fill(pubkey);
            privkeyBuf.fill(privkey);
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
