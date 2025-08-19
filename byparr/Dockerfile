FROM docker.io/fedora:latest

RUN dnf update -yy
RUN dnf install -yy python3-pip git xauth xorg-x11-server-Xvfb scrot curl chromium chromedriver python3-tkinter
RUN pip install uv

RUN cd /opt; git clone https://github.com/ThePhaseless/Byparr
WORKDIR /opt/Byparr
RUN mkdir -p /opt/Byparr/.cache
RUN uv sync
RUN chmod -R a+rw /opt/Byparr

EXPOSE 8191
ENTRYPOINT ["uv", "run", "main.py"]