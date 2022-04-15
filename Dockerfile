FROM node:12-alpine AS builder

WORKDIR /src

COPY . .
RUN yarn --pure-lockfile install

RUN yarn run build --spa

FROM rancher/rancher:v2.6.4
WORKDIR /var/lib/rancher
RUN rm -rf /usr/share/rancher/ui-dashboard/dashboard*
COPY --from=builder /src/dist /usr/share/rancher/ui-dashboard/dashboard