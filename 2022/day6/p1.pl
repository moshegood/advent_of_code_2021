use strict;
use warnings;

while (my $line = <>) {
    my @c = split('', $line);
    foreach my $i (3 .. length($line)-1) {
        my %seen;
        for my $j (0 .. 3) {
            last if $seen{$c[$i-$j]};
            $seen{$c[$i-$j]} = 1;
        }
        if (keys %seen == 4) {
            print $i+1 , "\n";
            last;
        }
    }
}

