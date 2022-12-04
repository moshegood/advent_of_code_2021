use strict;
use warnings;

my $sum = 0;
while (my $line = <>) {
    chomp $line;
    my @n = grep { /\d/ } $line =~ m/\d*/g;
    $sum++ if $n[0]<=$n[2] && $n[1]>=$n[3] || $n[0]>=$n[2] && $n[1]<=$n[3];
}

print "$sum\n";
