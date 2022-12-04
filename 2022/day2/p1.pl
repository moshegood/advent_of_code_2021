use strict;
use warnings;

my $sum = 0;
my %points = (
    2 => 0,
    1 => 6,
    0 => 3,
);


while (<>) {
    chomp;
    my ($p1,$p2)=split();
    my $me = ord($p2)-ord("X")+1;
    my $he = ord($p1)-ord("A")+1;
    $sum += $me + $points{($me-$he)%3};
}
print "$sum\n";
