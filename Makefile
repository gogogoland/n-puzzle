# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: croy <croy@student.42.fr>                  +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2016/03/15 16:41:50 by croy              #+#    #+#              #
#    Updated: 2016/03/21 18:13:39 by croy             ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

export GOPATH=$(PWD)

NAME = n-puzzle

B_PKG = go build
BIN = go install $(NAME)
RM = rm -rf bin/$(NAME)

PKGS = github.com/veandco/go-sdl2/sdl \
	   n-puzzle \
	   ui \
	   algo

all: $(NAME)

$(NAME):
	$(B_PKG) $(PKGS)
	$(BIN)

clean:
	$(RM)

fclean: clean

sub:
	$(shell git submodule init ; git submodule update)

re : fclean all
