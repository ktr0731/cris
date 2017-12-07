import APIClient from '../clients/api';

export default file => {
    return new APIClient().upload(file);
};
