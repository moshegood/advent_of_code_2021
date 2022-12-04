use strict;
use warnings;

my $sum = 0;

while (<>) {
    chomp;
    my ($p1,$p2)=split();
    my $me = ord($p2)-ord("X")+1;
    my $he = ord($p1)-ord("A")+1;
    $sum += 3*($me-1) + ($me+$he)%3 +1;
}
print "$sum\n";
