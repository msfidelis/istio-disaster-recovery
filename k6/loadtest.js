import http from 'k6/http';

export default  () => {
    const id = Math.random() +1 
    const url = 'http://orders-api.orders.svc.cluster.local:8080/orders/' + id;

    const payload = JSON.stringify({
    });

    const params = {
        headers: {
            'Content-Type': 'application/json'
        },
    };

    http.get(url, payload, params);
}