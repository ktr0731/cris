export default class {
    constructor() {}

    upload(file) {
        return fetch('http://localhost:8080/v1/files', {
            mode: 'cors',
            method: 'POST',
            body: file
        });
    }
}
