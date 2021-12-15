# A Study Sample of Kubernetes Admission Controller

see also: https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/

# Installation

```sh
$ kubectl apply -k deploy/
```
# A little test

```console
$ kubectl get po -l app=kube-admission
NAME                              READY   STATUS    RESTARTS   AGE
kube-admission-58c789595c-wth9j   1/1     Running   0          75s

$ kubectl scale deployment kube-admission --replicas=2
deployment.apps/kube-admission scaled

$ kubectl get po -l app=kube-admission -w
NAME                              READY   STATUS    RESTARTS   AGE
kube-admission-58c789595c-wth9j   1/1     Running   0          88s
kube-admission-58c789595c-wrlmx   0/1     Pending   0          5s
kube-admission-58c789595c-wrlmx   0/1     Pending   0          5s
kube-admission-58c789595c-wrlmx   0/1     ContainerCreating   0          5s
kube-admission-58c789595c-wrlmx   0/1     Running             0          8s
kube-admission-58c789595c-wrlmx   1/1     Running             0          8s

$ kubectl logs kube-admission-58c789595c-wth9j 
...
I1215 10:13:33.190637       1 main.go:37] handling request: {"kind":"AdmissionReview","apiVersion":"admission.k8s.io/v1"," ...

I1215 10:13:33.191145       1 main.go:87] always-allow-with-delay sleeping for 5 seconds
I1215 10:13:38.191770       1 main.go:89] calling always-allow
I1215 10:13:38.191904       1 main.go:68] sending response: &AdmissionReview{Request:nil,Response:&AdmissionResponse{UID:47a96663-03b2-45ec-8152-ee74a8441c7d,Allowed:true,Result:&v1.Status{ListMeta:ListMeta{SelfLink:,ResourceVersion:,Continue:,RemainingItemCount:nil,},Status:,Message:this webhook allows all requests,Reason:,Details:nil,Code:0,},Patch:nil,PatchType:nil,AuditAnnotations:map[string]string{},Warnings:[],},}
```
