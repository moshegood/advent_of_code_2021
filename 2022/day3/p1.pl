use strict;
use warnings;

my $sum = 0;
while (my $line = <>) {
    chomp $line;
    my $left = substr($line, 0, length($line)/2 );
    my $right = substr($line, length($line)/2);
    my %seen;
    foreach my $c (split('', $left)) {
        $seen{$c} = 1;
    }
    print "$line\n";
    foreach my $c (split('', $right)) {
        if ($seen{$c}) {
            print "Found $c in $line ($left, $right)\n";
            $sum += ord($c) - ord('a') + 1 if $c =~ m/[a-z]/;
            $sum += ord($c) - ord('A') + 27 if $c =~ m/[A-Z]/;
            last;
        }
    }
}

print "$sum\n";
