import 'riot';

export default class {
    constructor() {
        this.uploaded = JSON.parse(localStorage['uploaded'] || '[]');

        riot.observable(this);
    }

    saveAccount(pubkey) {}

    setAccount(account) {
        this.account = account;
        this.trigger('get_account', account);
    }

    addUploadedFile(file) {
        this.uploaded.push(file);
        this.trigger('add_uploaded_file', file);
    }
}
