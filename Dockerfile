FROM node:12-alpine AS builder

WORKDIR /src

COPY . .
RUN yarn --pure-lockfile install

ENV ROUTER_BASE="/dashboard"
RUN yarn run build --spa

FROM <REPOSITORY>:<VERSION>
WORKDIR /var/lib/rancher
RUN rm -rf /usr/share/rancher/ui-dashboard/dashboard*
COPY --from=builder /src/dist /usr/share/rancher/ui-dashboard/dashboard
