import 'riot';

import '../tags/app.tag';
import '../tags/upload.tag';
import '../tags/status.tag';

import 'riot-mui';

import Store from './stores/store';
import APIClient from './clients/api';
import EthClient from './clients/ethereum';

riot.mount('*', {
    store: new Store(),
    apiClient: new APIClient(),
    ethClient: new EthClient()
});
