/* Base case when Input is empty */
reachable(Nfa, StartState, FinalState, []) :- StartState = FinalState.

/* Set States to be next possible steps from Start State using Input */
/* Recurse reachable with States */
reachable(Nfa, StartState, FinalState, [Input|Rest]) :- 
    transition(Nfa, StartState, Input, States),
    reachable(Nfa, States, FinalState, Rest).

/* When StartState is a list, recurse on head of list and rest of list */
/* Return true if either is true */
reachable(Nfa, [H|T], FinalState, Input) :- 
    reachable(Nfa, H, FinalState, Input) ; reachable(Nfa, T, FinalState, Input).