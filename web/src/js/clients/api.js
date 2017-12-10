import { Base64 } from 'js-base64';

export default class {
    constructor() {}

    upload(content, signature) {
        const payload = {
            content: Base64.btoa(content.toString('latin1')),
            signature: Base64.btoa(signature.toString('latin1')),
            pubkey: localStorage['pubkey']
        };
        return fetch('http://localhost:8080/v1/files', {
            mode: 'cors',
            method: 'POST',
            body: JSON.stringify(payload)
        });
    }
}
