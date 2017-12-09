import 'riot';

import '../tags/app.tag';
import '../tags/upload.tag';
import '../tags/status.tag';

import 'riot-mui';

import 'clipboard';

import Store from './stores/store';
import APIClient from './clients/api';
import EthClient from './clients/ethereum';
import share from './share';

(() => {
    const parseQuery = queryString => {
        if (queryString.length == 0) {
            return {};
        }

        let queries = {};
        queryString
            .slice(1)
            .split('&')
            .forEach(pair => {
                const s = pair.split('=');
                queries[s[0]] = s.length == 2 ? s[1] : '';
            });

        return queries;
    };

    const query = parseQuery(document.location.search);
    if (query['url']) {
        share(new EthClient(), query['url']);
        return;
    }

    riot.mount('*', {
        store: new Store(),
        apiClient: new APIClient(),
        ethClient: new EthClient()
    });
})();
