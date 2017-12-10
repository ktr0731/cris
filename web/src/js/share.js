import { Base64 } from 'js-base64';

const parseURL = url => {
    return Base64.decode(url).split('.');
};
const fetchContent = url => {
    const headers = new Headers();
    headers.append(
        'Accept',
        'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8'
    );
    return fetch(url, {
        mode: 'cors',
        headers: headers
    }).then(res => res.blob());
};

const verifyFileHash = (ethClient, hash) => {
    return ethClient.fetch('', hash);
};

const doDownload = content => {
    const url = URL.createObjectURL(content);
    const body = document.querySelector('body');
    const e = document.createElement('a');
    e.setAttribute('href', url);
    e.setAttribute('download', '');
    body.appendChild(e);
    e.click();
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
        .then(doDownload)
        .catch(e => {
            console.log(e);
        });
};
