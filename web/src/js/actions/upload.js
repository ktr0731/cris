import APIClient from '../clients/api';
import sjcl from 'sjcl';
import ed25519 from 'supercop.js';

export default (file, store) => {
    // TODO: send signature also
    return new Promise((resolve, reject) => {
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
            console.log(
                'HASH:',
                sjcl.codec.hex.fromBits(sjcl.hash.sha256.hash(encryptedContent))
            );

            resolve(encryptedContent);
        };
        reader.readAsText(file);
    }).then(new APIClient().upload);
};
