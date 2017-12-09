import APIClient from '../clients/api';
import EthClient from '../clients/ethereum';
import ed25519 from 'supercop.js';
import { Base64 } from 'js-base64';
import ab2buf from 'arraybuffer-to-buffer';
import { sha3_256 } from 'js-sha3';

export default (file, store) => {
    // TODO: send signature also
    let signature = null;
    let token = null;
    return new Promise((resolve, _) => {
        const reader = new FileReader();
        reader.onload = e => {
            const pubkey = Buffer.from(
                Base64.atob(localStorage['pubkey']),
                'latin1'
            );
            const privkey = Buffer.from(
                Base64.atob(localStorage['privkey']),
                'latin1'
            );
            const contentBuf = ab2buf(reader.result);
            // const pubkeyBuf = new Buffer(pubkey.length);
            // const privkeyBuf = new Buffer(privkey.length);
            // pubkeyBuf.fill(pubkey);
            // privkeyBuf.fill(privkey);
            // signature = ed25519.sign(contentBuf, pubkeyBuf, privkeyBuf);
            signature = ed25519.sign(contentBuf, pubkey, privkey);

            console.log(
                pubkey.toString('hex'),
                sha3_256(Buffer.from(contentBuf).toString())
            );

            return resolve(contentBuf);
        };
        reader.readAsArrayBuffer(file);
    })
        .then(content => new APIClient().upload(content, signature))
        .then(res => res.json())
        .then(res => {
            token = res.token;
            return new EthClient().store(hash);
        })
        .then(txHash => {
            console.log('stored:', txHash);
            store.addUploadedFile({
                name: file.name,
                hash: hash,
                token: token,
                txHash: txHash,
                date: file.lastModifiedDate
            });
        });
};
