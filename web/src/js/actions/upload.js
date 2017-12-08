import APIClient from '../clients/api';
import EthClient from '../clients/ethereum';
import ed25519 from 'supercop.js';

export default (file, store) => {
    // TODO: send signature also
    let hash = null;
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
            console.log('FOO');
            store.addUploadedFile({
                name: file.name,
                hash: hash,
                token: res.token,
                date: file.lastModifiedDate
            });
        });
    // .then(() => {
    //     console.log('FOO');
    //     return new EthClient().store(hash);
    // });
};
