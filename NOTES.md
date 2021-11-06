
## K6 run

```bash
kubectl run -i --tty clients-02 --image=fidelissauro/k6-stack -n chip -- bash
k6 run --vus 20 --iterations 100 loadtest.js
k6 run --vus 20 --duration 60s loadtest.js
```

