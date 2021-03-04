
% A list is a 1-D array of numbers.
% A matrix is a 2-D array of numbers, stored in row-major order.

% You may define helper functions here.
adjacent([A,B|_], A, B).

adjacent([_|T], A, B) :-
    adjacent(T, A, B).

transpose_top_row([], [], []).
transpose_top_row([[H|T]|RestRows], [H|RestH], [T|RestT]) :- 
    transpose_top_row(RestRows, RestH, RestT).

neighbors([], _, _).
neighbors([FirstRow|RestRows], A, B) :- 
    are_adjacent(FirstRow, A, B); neighbors(RestRows, A, B).

% are_adjacent(List, A, B) returns true iff A and B are neighbors in List.
are_adjacent(List, A, B) :-
    adjacent(List, A, B); adjacent(List, B, A).
    
% transpose(Matrix, Answer) returns true iff Answer is the transpose of the 2D
% matrix Matrix
transpose([], []).
transpose([[]|_], []).
transpose(Matrix, [FirstRow|RestRows]) :- 
    transpose_top_row(Matrix, FirstRow, RestMatrix), transpose(RestMatrix, RestRows).

% are_neighbors(Matrix, A, B) returns true iff A and B are neighbors in the 2D
% matrix Matrix.
are_neighbors(Matrix, A, B) :-
    neighbors(Matrix, A, B); neighbors(Transpose, A, B), transpose(Matrix, Transpose).

    

