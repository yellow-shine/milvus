# Copyright (C) 2019-2022 Zilliz. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under the License.

FROM rockylinux/rockylinux:9.3

ARG TARGETARCH


RUN dnf install -y  git wget
RUN dnf install -y epel-release
RUN dnf groupinstall -y "Development Tools"
RUN dnf install -y wget curl which \
  git make automake python3-devel  python3-pip \
  gcc gcc-c++ gcc-gfortran \
  clang llvm-devel \
  libaio libuuid-devel zip unzip \
  ccache lcov libtool m4 autoconf automake
RUN rm -rf /var/cache/yum/*


RUN pip3 install conan==1.61.0
RUN dnf install -y perl

RUN echo "target arch $TARGETARCH"
RUN wget -qO- "https://cmake.org/files/v3.28/cmake-3.28.1-linux-x86_64.tar.gz" | tar --strip-components=1 -xz -C /usr/local

RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
RUN tar -xvf go1.21.6.linux-amd64.tar.gz -C /usr/local
RUN echo "export PATH=/usr/local/bin:$PATH:/usr/local/go/bin" >> ~/.bashrc
RUN source ~/.bashrc    

COPY --chown=0:0 build/docker/builder/entrypoint.sh /

RUN curl https://sh.rustup.rs -sSf | \
    sh -s -- --default-toolchain=1.73 -y

ENV PATH=/root/.cargo/bin:$PATH

ENTRYPOINT [ "/entrypoint.sh" ]
CMD ["tail", "-f", "/dev/null"]
