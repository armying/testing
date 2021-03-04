mymember(X, [X|_]).
mymember(X, [_|Y]) :- mymember(X, Y).

common_members([X|_], List) :-
    mymember(X, List).

common_members([_|Y], List) :-
    common_members(Y, List).

/* All novels published either during the year 1953 or during the year 1996*/
year_1953_1996_novels(Book) :-
    novel(Book, 1953); novel(Book, 1996).

/* List of all novels published during the period 1800 to 1900 (not inclusive)*/
period_1800_1900_novels(Book) :-
    novel(Book, Year), (Year > 1800, Year < 1900).

/* Characters who are fans of LOTR */
lotr_fans(Fan) :-
    fan(Fan, List), mymember(the_lord_of_the_rings, List).

/* Authors of the novels that heckles is fan of. */
heckles_idols(Author) :-
    author(Author, Authorlist), fan(heckles, Fanlist), common_members(Authorlist, Fanlist).

/* Characters who are fans of any of Robert Heinlein's novels */
heinlein_fans(Fan) :-
    fan(Fan, Fanlist), author(robert_heinlein, Authorlist), common_members(Fanlist, Authorlist).

/* Novels common between either of Phoebe, Ross, and Monica */
mutual_novels(Book) :-
    fan(phoebe, Phoebelist), fan(ross, Rosslist), fan(monica, Monicalist), 
    ((mymember(Book, Phoebelist), mymember(Book, Rosslist)); (mymember(Book, Phoebelist), mymember(Book, Monicalist)); (mymember(Book, Rosslist), mymember(Book, Monicalist))).
   