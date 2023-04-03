import { useState, useCallback, useRef, useEffect } from 'react';

import { fetchEventSource } from '@microsoft/fetch-event-source';

const URL = process.env.REACT_APP_API_ENDPOINT || '';

const agentSSE = async (url, method, headers, body, setOk, setData, signal) => {
  await fetchEventSource(url, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
    signal,
    onopen(response) {
      setOk(response.ok && response.status === 200);
      if (!response.ok) throw new Error(response);
    },
    onmessage(event) {
      const parsedData = JSON.parse(event.data);
      setData(parsedData);
    },
    onerror(err) {
      throw new Error(err);
    },
  });
};

const agent = async (url, method, headers, body, setOk, setData) => {
  const response = await fetch(url, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
  });

  setOk(response.ok);
  if (!response.ok) throw response;

  if (method === 'GET' || (url.includes('/api/auth'))) {
    const json = await response.json();
    setData(json);
  }
};

const parseParams = (rawurl, params) => {
  let url = rawurl;
  if (params) {
    const paramsArr = Object.keys(params);

    paramsArr.forEach((p) => {
      if (url.includes(p)) {
        url = url.replace(`{${p}}`, params[p]);
      }
    });
  }

  return url;
};

const parseQuery = (query) => {
  let queryParams = '';

  if (query) {
    const [querySymbol, addSymbol, equalSymbol] = ['?', '&', '='];
    const queryArr = Object.keys(query);
    let currParam = 1;

    queryParams += querySymbol;
    queryArr.forEach((q) => {
      queryParams = queryParams + q + equalSymbol + query[q];

      if (currParam < queryArr.length) {
        queryParams += addSymbol;
        currParam += 1;
      }
    });
  }

  return queryParams;
};

const Service = (rawURL, method = 'GET', sse = false) => {
  const [ok, setOk] = useState(false);
  const [data, setData] = useState();
  const [error, setError] = useState();
  const [loading, setLoading] = useState(false);

  const ctrlRef = useRef(null);

  useEffect(() => {
    ctrlRef.current = new AbortController();
    return () => (ctrlRef.current.abort());
  }, [ctrlRef]);

  const request = useCallback(async ({ params, query, body } = {}) => {
    setLoading(true);
    let url = parseParams(rawURL, params);
    url += parseQuery(query);

    const headers = {};
    headers.Timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;

    if (body) {
      headers['Content-Type'] = 'application/json';
    }

    try {
      if (!sse) {
        await agent(url, method, headers, body, setOk, setData);
      } else {
        await agentSSE(url, method, headers, body, setOk, setData, ctrlRef.current.signal);
      }
    } catch (err) {
      const errorStatus = err.status || 0;
      const errorMessage = err.statusText || err.message || 'Unexpected Error!';

      setError({ status: errorStatus, message: errorMessage });
    } finally {
      setLoading(false);
    }
  }, [rawURL, method, sse]);

  return {
    ok,
    data,
    error,
    loading,
    request,
  };
};

const requests = {

  del: (url) => Service(url, 'DELETE'),

  get: (url) => Service(url),

  post: (url) => Service(url, 'POST'),

  put: (url) => Service(url, 'PUT'),

  getSSE: (url) => Service(url, 'GET', true),

};

const games = {
  basePath: '/api/pub/games',

  useAll: () => {
    const url = (games.basePath).startsWith('/') ? URL + games.basePath : games.basePath;

    return requests.get(url);
  },

  useCalendarSSE: () => {
    const path = `${games.basePath}/{gid}/calendar/sse`;
    const url = path.startsWith('/') ? URL + path : path;

    return requests.getSSE(url);
  },

};

const bookings = {
  basePath: '/api/pub/games/{gid}/bookings',

  useCreate: () => {
    const url = (bookings.basePath).startsWith('/') ? URL + bookings.basePath : bookings.basePath;

    return requests.post(url);
  },
};

export default {
  games,
  bookings,
};
