FROM node:12-alpine AS builder

WORKDIR /root/dashboard

COPY . .
RUN yarn --pure-lockfile install

RUN yarn run build --spa

FROM rancher/rancher:v2.6.4
WORKDIR /var/lib/rancher
RUN rm -rf /usr/share/rancher/ui-dashboard/dashboard*
COPY --from=builder /root/dashboard/dist /usr/share/rancher/ui-dashboard/dashboard