import 'riot';

export default class {
    constructor() {
        riot.observable(this);
    }

    saveAccount(pubkey) {}

    setUploadedFiles() {
        console.log('set uploaded');
        this.uploaded = JSON.parse(localStorage['uploaded'] || '[]');
        this.trigger('set_uploaded_files', this.uploaded);
    }

    setAccount(account) {
        this.account = account;
        this.trigger('get_account', account);
    }

    addUploadedFile(file) {
        this.uploaded.push(file);
        this.save();

        this.trigger('add_uploaded_file', file);
    }

    save() {
        localStorage['uploaded'] = JSON.stringify(this.uploaded);
    }
}
