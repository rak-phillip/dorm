FROM node:lts AS builder

WORKDIR /home/dashboard

COPY . .
RUN yarn --pure-lockfile install

RUN yarn run build --spa

FROM rancher/rancher:v2.6.4
WORKDIR /var/lib/rancher
RUN rm -rf /usr/share/rancher/ui-dashboard/dashboard*
COPY --from=builder /home/dashboard/dist /usr/share/rancher/ui-dashboard/dashboard