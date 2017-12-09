import { Base64 } from 'js-base64';
import ed25519 from 'supercop.js';

const parseURL = url => {
    return Base64.decode(url).split('.');
};
const fetchContent = url => {
    return fetch(url, { mode: 'cors' }).then(res => res.blob());
};

const verifyFileHash = (ethClient, hash) => {
    return ethClient.fetch('', hash);
};

const decodeContent = encrypted => {
    //ed25519.veri
};

export default (ethClient, url) => {
    const [token, hash, privkey] = parseURL(url);

    const actualURL = `http://localhost:8080/v1/files/${token}`;

    verifyFileHash(ethClient, hash)
        .then(verified => {
            if (!verified) {
                throw Error('tempering detected');
            }
        })
        .then(() => fetchContent(actualURL))
        .then(decodeContent)
        .then(content => {
            console.log(content);
        })
        .catch(e => {
            console.log(e);
        });
};
