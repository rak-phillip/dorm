FROM node:lts AS builder

WORKDIR /run/dashboard

COPY . .
RUN yarn --pure-lockfile install

ENV ROUTER_BASE="/dashboard"
RUN yarn run build --spa

FROM rancher/rancher:v2.6.4
WORKDIR /var/lib/rancher
RUN rm -rf /usr/share/rancher/ui-dashboard/dashboard*
COPY --from=builder /run/dashboard/dist /usr/share/rancher/ui-dashboard/dashboard