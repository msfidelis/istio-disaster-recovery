import http from 'k6/http';

export default  () => {
    const id = Math.random() +1 
    const url = 'http://api.orders.k8s.cluster/orders/' + id;

    const payload = JSON.stringify({
    });

    const params = {
        headers: {
            'Content-Type': 'application/json'
        },
    };

    http.get(url, payload, params);
}